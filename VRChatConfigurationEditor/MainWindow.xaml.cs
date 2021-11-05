using Microsoft.Web.WebView2.Core;
using Microsoft.Web.WebView2.Wpf;
using System;
using System.Diagnostics;
using System.IO;
using System.Reflection;
using System.Windows;
using System.Windows.Controls;

namespace VRChatConfigurationEditor
{
    /// <summary>
    /// MainWindow.xaml 的交互逻辑
    /// </summary>
    public partial class MainWindow : Window
    {
        private readonly WebView2 _webView = new WebView2();
        private Bind _bind;
        private const string VirtualHost = "vrc.conf.editor.invalid";
        public MainWindow()
        {
            InitializeComponent();
            AppInit();
            WebViewInit();
        }

        private void AppInit()
        {
            string systemLanguage = System.Threading.Thread.CurrentThread.CurrentCulture.Name;
            ResourceDictionary lang = LoadLanguage(systemLanguage);
            if (lang == null && systemLanguage.Contains("-")) lang = LoadLanguage(systemLanguage.Split('-')[0]);
            if (lang == null) return;
            if (Resources.MergedDictionaries.Count > 0) Resources.MergedDictionaries.Clear();
            Resources.MergedDictionaries.Add(lang);
        }

        private void WebViewInit()
        {
            if (GetWebView2Version() == "") return;

            _webView.CoreWebView2InitializationCompleted += _webView_CoreWebView2InitializationCompleted;
            _webView.Loaded += _webView_Loaded;

            DockPanel panel = new DockPanel();
            panel.Children.Add(_webView);
            Content = panel;
        }

        private void _webView_CoreWebView2InitializationCompleted(object sender, CoreWebView2InitializationCompletedEventArgs e)
        {
            if (!e.IsSuccess) return;
            CoreWebView2Settings settings = _webView.CoreWebView2.Settings;
#if !DEBUG
            settings.AreDevToolsEnabled = false;
#endif
            settings.IsPinchZoomEnabled = false;
            settings.IsZoomControlEnabled = false;
            settings.IsPasswordAutosaveEnabled = false;
            settings.IsStatusBarEnabled = false;
        }

        private void CoreWebView2_WebMessageReceived(object sender, CoreWebView2WebMessageReceivedEventArgs args)
        {
            _bind.MessageReceived(args);
        }

        private async void _webView_Loaded(object sender, RoutedEventArgs e)
        {
            AssemblyName name = Assembly.GetExecutingAssembly().GetName();
            string cacheFolderPath = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.LocalApplicationData), name.Name);
            CoreWebView2Environment env = await CoreWebView2Environment.
                CreateAsync(null, cacheFolderPath).ConfigureAwait(true);

            await _webView.EnsureCoreWebView2Async(env);

            _webView.CoreWebView2.SetVirtualHostNameToFolderMapping(
                VirtualHost, @"public", CoreWebView2HostResourceAccessKind.DenyCors);

            _webView.Source = new Uri($"https://{VirtualHost}/index.html");
            //_webView.Source = new Uri("http://localhost:3000/index.html");
            _webView.CoreWebView2.WebMessageReceived += CoreWebView2_WebMessageReceived;

            _bind = new Bind(_webView);
        }

        private void RuntimeDownloadButton_OnClick(object sender, RoutedEventArgs e)
        {
            //string url = "https://go.microsoft.com/fwlink/p/?LinkId=2124703";
            string url = "https://developer.microsoft.com/microsoft-edge/webview2/#download-section";
            _ = Process.Start(new ProcessStartInfo(url));
        }

        private void Window_Closing(object sender, System.ComponentModel.CancelEventArgs e)
        {
            Visibility = Visibility.Hidden;
            _webView.Dispose();
        }

        private void Window_Initialized(object sender, EventArgs e)
        {
            Version version = Assembly.GetExecutingAssembly().GetName().Version;
            string buildDate = File.GetLastWriteTime(Assembly.GetExecutingAssembly().Location).ToString("yyyyMMdd");
            Title = $"{Title} v{version.Major}.{version.Minor}.{version.Build} build-{buildDate}";
        }
    }
}
