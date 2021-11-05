using Microsoft.Web.WebView2.Core;
using System;
using System.Windows;

namespace VRChatConfigurationEditor
{
    public partial class MainWindow
    {
        public static string GetWebView2Version()
        {
            try { return CoreWebView2Environment.GetAvailableBrowserVersionString(); }
            catch (Exception) { return ""; }
        }

        public ResourceDictionary LoadLanguage(string name)
        {
            try
            {
                return Application.LoadComponent(new Uri($@"languages\{name}.xaml", UriKind.Relative)) as
                    ResourceDictionary;
            }
            catch (Exception) { return null; }
        }
    }

}
