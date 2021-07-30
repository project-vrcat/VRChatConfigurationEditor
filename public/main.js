new Vue({
  el: '#app',
  data: {
    disabled: true,
    removeCache: false,
    config_file: '',
    vrchat_dir: '',
    old_config: {},
    config: {
      disableRichPresence: false,
      cache_directory: '',
      cache_size: 20,
      cache_expiry_delay: 30,
      camera_res_height: 1080,
      camera_res_width: 1920,
      screenshot_res_height: 1080,
      screenshot_res_width: 1920,
    },
    cameraRes: '1080',
    screenshotRes: '1080',
  },
  watch: {
    cameraRes(value) {
      switch (value) {
        case '4K':
          this.config.camera_res_height = 2160;
          this.config.camera_res_width = 3840;
          break;
        case '2K':
          this.config.camera_res_height = 1440;
          this.config.camera_res_width = 2560;
          break;
        case '720':
          this.config.camera_res_height = 720;
          this.config.camera_res_width = 1280;
          break;
        default:
          this.config.camera_res_height = 1080;
          this.config.camera_res_width = 1920;
          break;
      }
    },
    screenshotRes(value) {
      switch (value) {
        case '4K':
          this.config.screenshot_res_height = 2160;
          this.config.screenshot_res_width = 3840;
          break;
        case '2K':
          this.config.screenshot_res_height = 1440;
          this.config.screenshot_res_width = 2560;
          break;
        case '720':
          this.config.screenshot_res_height = 720;
          this.config.screenshot_res_width = 1280;
          break;
        default:
          this.config.screenshot_res_height = 1080;
          this.config.screenshot_res_width = 1920;
          break;
      }
    }
  },
  created() {
    checkUpdate().then(update => {
      if (update && confirm("New update available.\nPress the \"OK\" button to download.") === true) open("https://lumina.moe/downloads")
    })
    appVersion().then(version => {
      document.title += " " + version
    })
    vrchatPath().then(dir => {
      this.vrchat_dir = dir
      this.config_file = dir + "\\config.json";
      this.load_config()
    }).catch(err => {
      alert(err)
    })
  },
  methods: {
    load_config() {
      readTextFile(this.config_file).then(data => {
        let config = JSON.parse(data)
        if (config.disableRichPresence === undefined) config.disableRichPresence = false
        if (config.cache_directory === undefined) config.cache_directory = ""
        if (config.cache_size === undefined) config.cache_size = 20
        if (config.cache_expiry_delay === undefined) config.cache_expiry_delay = 30
        switch (config.screenshot_res_height) {
          case 2160:
            this.screenshotRes = '4K'
            break;
          case 1440:
            this.screenshotRes = '2K'
            break;
          case 720:
            this.screenshotRes = '720'
            break;
          default:
            this.screenshotRes = '1080'
            break;
        }
        switch (config.camera_res_height) {
          case 2160:
            this.cameraRes = '4K'
            break;
          case 1440:
            this.cameraRes = '2K'
            break;
          case 720:
            this.cameraRes = '720'
            break;
          default:
            this.cameraRes = '1080'
            break;
        }
        this.old_config = Object.assign({}, config)
        this.config = config
        this.disabled = false
      }).catch(err => {
        this.disabled = false
        console.error(err)
      })
    },
    select() {
      selectDirectory("Custom Cache Directory Location").then(dir => {
        if (dir) {
          this.config.cache_directory = dir
        }
      }).catch(err => {
        console.error(err)
      })
    },
    save() {
      writeTextFile(this.config_file, JSON.stringify(this.config)).then(_ => {
        alert("success")
      }).catch(err => {
        alert(err)
      })
    },
    reset() {
      this.config = Object.assign({}, this.old_config);
    },
    removeAllCache() {
      this.removeCache = true
      removeAll(this.vrchat_dir + "\\Cache-WindowsPlayer").then(_ => {
        this.removeCache = false
      }).catch(err => {
        alert(err)
        this.removeCache = false
      })
    }
  },
})