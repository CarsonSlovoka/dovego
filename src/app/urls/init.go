package urls

func InitURLs() {
    initFavicon()
    initPlugin()
    initSystemURL()
    initSystemStaticResource() // 這個會對整個 "/" 進行過濾，因此要放在最後執行。如果router有很多重複，只會執行最先建立的那一個
}
