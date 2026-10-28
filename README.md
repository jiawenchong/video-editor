# 馬卡龍剪輯室

藍米色加馬卡龍配色的簡易影片剪輯器。影片只在你自己的電腦裡處理，不會上傳到任何地方，也不需要網路。

## 怎麼用

**方法一：exe（推薦）**
到 [Releases](https://github.com/jiawenchong/video-editor/releases) 下載 `MacaronCut.exe`，點兩下就會用 Edge（或 Chrome）開成一個獨立的剪輯視窗。

> 第一次執行如果出現「Windows 已保護您的電腦」，按「其他資訊 → 仍要執行」。這是因為 exe 沒有數位簽章，不是病毒。若公司電腦完全不能執行 exe，請改用方法二。

**方法二：直接開 HTML**
下載 `index.html`，用 Edge 或 Chrome 打開就能用，功能完全一樣。

## 功能

| 功能 | 按鈕 | 快捷鍵 |
|---|---|---|
| 播放 / 暫停 | ▶ | 空白鍵 |
| 在播放頭分割 | 分割 | `S` |
| 刪掉播放頭左邊（剪頭） | 剪頭 | `Q` |
| 刪掉播放頭右邊（剪尾） | 剪尾 | `W` |
| 刪除選取的片段（剪中間：先分割兩刀再刪） | 刪除片段 | `Delete` |
| 調整片段速度 0.25×–8×（聲音不會變調） | 右側速度 | — |
| 片段前後移動 | 往前 / 往後 | — |
| 復原 / 重做 | 復原 / 重做 | `Ctrl+Z` / `Ctrl+Shift+Z` |
| 逐格移動 | — | `←` `→`（加 Shift 移 1 秒） |

可以一次加入多支影片，會依序接在一起。

## 匯出格式怎麼選

- **MP4 · H.264 + 標準畫質**：解析度不變、檔案小，手機、LINE、Teams、PowerPoint 都能播。日常選這個。
- **MP4 · H.265**：同畫質再小約 40%，但舊電腦或部分網頁可能播不了。只有電腦支援時才能選。
- **精簡**：檔案最小，適合傳訊息；畫面細節會少一點。

匯出的檔案會存到瀏覽器的「下載」資料夾。

## 需求

Windows 10/11，最新版 Edge 或 Chrome（需要 WebCodecs）。iPhone 拍的 HEVC 影片若打不開，請先轉成 H.264。

## 自己重新編譯 exe

```bash
GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "-H=windowsgui -s -w" -o MacaronCut.exe .
```

`index.html` 會被打包進 exe 裡，改完網頁重新編譯即可。

## 授權

匯出 MP4 使用 [mp4-muxer](https://github.com/Vanilagy/mp4-muxer)（MIT License），已內嵌在 `index.html`。
