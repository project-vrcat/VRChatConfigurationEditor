using Microsoft.Web.WebView2.Core;
using Microsoft.Web.WebView2.Wpf;
using Newtonsoft.Json;
using Ookii.Dialogs.Wpf;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Runtime.InteropServices;

namespace VRChatConfigurationEditor
{
    class Message
    {
        public string Method { get; set; }
        public List<string> Args { get; set; }
    }

    class Bind
    {
        private readonly WebView2 _webView;
        private readonly string _vrcPath;

        public Bind(WebView2 webView)
        {
            _webView = webView;

            string localLow = GetLocalLow();
            if (localLow != "") _vrcPath = Path.Combine(localLow, @"VRChat\VRChat");
        }

        public void MessageReceived(CoreWebView2WebMessageReceivedEventArgs args)
        {
            string jsonMessage = args.WebMessageAsJson;
            Message message = JsonConvert.DeserializeObject<Message>(jsonMessage);
            if (message == null) return;
            switch (message.Method)
            {
                case "openURL":
                    if (message.Args == null || message.Args.Count < 1) break;
                    OpenURL(message.Args[0]);
                    break;
                case "selectDirectory":
                    if (message.Args == null || message.Args.Count < 1) return;
                    SelectDirectory(message.Args[0], message.Args.Count >= 2 ? message.Args[1] : "");
                    break;
                case "loadConfigFile":
                    if (message.Args == null || message.Args.Count < 1) return;
                    LoadConfigFile(message.Args[0]);
                    break;
                case "saveConfigFile":
                    if (message.Args == null || message.Args.Count < 1) return;
                    SaveConfigFile(message.Args[0], message.Args[1]);
                    break;
            }
        }

        private void OpenURL(string url)
        {
            Process.Start(new ProcessStartInfo(url));
        }

        private void SelectDirectory(string key, string title)
        {
            VistaFolderBrowserDialog dialog = new VistaFolderBrowserDialog()
            {
                Description = title,
                UseDescriptionForTitle = true
            };
            var ok = dialog.ShowDialog();
            string path;
            if (!ok.HasValue || !(bool)ok) path = "";
            else path = dialog.SelectedPath;
            string data = JsonConvert.SerializeObject(new Message()
            {
                Method = "selectDirectory",
                Args = new List<string>() { key, path }
            });
            _webView.CoreWebView2.PostWebMessageAsJson(data);
        }

        [DllImport("shell32.dll")]
        static extern int SHGetKnownFolderPath([MarshalAs(UnmanagedType.LPStruct)] Guid rfid, uint dwFlags, IntPtr hToken, out IntPtr pszPath);
        public string GetLocalLow()
        {
            // https://docs.microsoft.com/en-us/windows/win32/shell/knownfolderid#FOLDERID_LOCALAPPDATALOW
            Guid localLowId = new Guid("A520A1A4-1780-4FF6-BD18-167343C5AF16");
            int hr = SHGetKnownFolderPath(localLowId, 0, IntPtr.Zero, out IntPtr pPath);
            return hr >= 0 ? Marshal.PtrToStringAuto(pPath) : "";
        }

        public void LoadConfigFile(string key)
        {
            string configFile = Path.Combine(_vrcPath, "config.json");
            string configData = "{}";
            if (File.Exists(configFile)) configData = File.ReadAllText(configFile);
            string data = JsonConvert.SerializeObject(new Message()
            {
                Method = "loadConfigFile",
                Args = new List<string>() { key, configData }
            });
            _webView.CoreWebView2.PostWebMessageAsJson(data);
        }

        public void SaveConfigFile(string key, string data)
        {
            string configFile = Path.Combine(_vrcPath, "config.json");
            Message message = new Message()
            {
                Method = "saveConfigFile",
                Args = new List<string>() { key, "success" }
            };
            try
            {
                File.WriteAllText(configFile, data);
            }
            catch (Exception e)
            {
                Debug.WriteLine(e);
                message.Args[1] = e.Message;
            }
            _webView.CoreWebView2.PostWebMessageAsJson(JsonConvert.SerializeObject(message));
        }
    }
}
