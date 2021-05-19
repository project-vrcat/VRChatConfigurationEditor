new Vue({
    el: '#app',
    data: {
        disabled: true,
        removeCache: false,
        config_file: '',
        vrchat_dir: '',
        old_config: {},
        config: {
            cache_directory: '',
        }
    },
    created() {
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
                if (config.cache_directory === undefined) config.cache_directory = ""
                this.old_config = Object.assign({}, config)
                this.config = config
                this.disabled = false
            }).catch(err => {
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