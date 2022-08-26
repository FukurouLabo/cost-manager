# Install
wailsのインストールを行います。
```bash
go get -u github.com/wailsapp/wails/cmd/wails
```

# Setup
wailsのセットアップを行います。  
ユーザ名とメールアドレスを設定します。(後から変更可能です)
```bash
wails setup
```

# Init
wailsの初期設定を行います。  
プロジェクト名を入力するとフロントエンドのフレームワークを指定します。  
今回は6の`Vue3 JS`を選択します。
```bash
wails init
```

# Build
初期設定が完了するとサンプルプロジェクトが生成されるので、試しにビルドします。  
ビルドが完了するとbuild/ディレクトリに実行形式バイナリファイルが生成されます。
```bash
wails build
```
