# yalzo
![travis ci](https://travis-ci.org/mizkei/yalzo.svg)

コマンドラインベースのシンプルTODO管理

# Install

```
$ go get github.com/mizkei/yalzo/cmd/yal
```

binary coming soon...

## マニュアル

### モード切替
操作 | キーバインド
--- | ---
ノーマル | ESC or C-[
挿入 | C-w
入れ替え | C-x
Todoのラベル変更 | C-l
ラベル追加 | C-v
ラベル削除 | C-c

### ノーマルモード時のキーバインド
操作 | キーバインド
--- | ---
選択 | Space
カーソルを1つ下に移動 | j
カーソルを5つ下に移動 | J
カーソルを末尾に移動 | G
カーソルを1つ上に移動 | k
カーソルを5つ上に移動 | K
Title変更 | C-r
Label変更 | C-l
Archive/Todo切り替え| (選択 =>) C-a
Todoの削除 | 選択 => C-d
順番入れ替え | C-x => 移動 => Enter

### ラベル削除モード時のキーバインド
操作 | キーバインド
--- | ---
カーソルを1つ下に移動 | j
カーソルを1つ上に移動 | k
ラベル削除 | C-d

# TODO

- ラベル編集モードの追加
- ページング機能の追加
