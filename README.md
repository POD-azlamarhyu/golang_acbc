# golang_acbc - Go言語アルゴリズム練習プロジェクト

Go言語のアルゴリズムやデータ構造の学習・練習用プロジェクトです。
Makefileを使用して一般的なタスクを簡単に実行できます。

## 使用方法

### ビルド関連

#### `make build` - ビルドと実行
```bash
make build
```
Go のバイナリをビルドして即座に実行します。
- 実行内容: `go build ./main.go && ./main i`

#### `make run` - 直接実行
```bash
make run
```
ビルドせずに直接プログラムを実行します（開発時に便利）。
- 実行内容: `go run .`

#### `make ts` - テスト実行
```bash
make ts
```
テストを実行します。

### ファイル管理関連

プログラミングコンテストの解いたファイルを整理するコマンド：

#### `make mv` - 解いたファイルを移動
```bash
make mv
```
解いたファイルを指定のディレクトリに移動します。

#### `make mvc` - コンテスト問題をコンテストフォルダに移動
```bash
make mvc
```
コンテスト関連の解いたファイルをコンテスト用フォルダに整理します。

#### `make mbc` - ABC問題（AtCoder Beginner Contest）をABCフォルダに移動
```bash
make mvbc
```
AtCoder Beginner Contestの解いたファイルを専用フォルダに整理します。

### その他

#### `make tml` - テンプレート生成
```bash
make tml
```
新しいプログラムファイル用のテンプレートを生成します。

## ディレクトリ構成

```
.
├── main.go                    # メインプログラム
├── Makefile                   # タスク定義ファイル
├── README.md                  # このファイル
├── go.mod                     # Go モジュール定義
└── src/
    └── script/                # 実行用シェルスクリプト
        ├── build.sh           # ビルドスクリプト
        ├── run.sh             # 実行スクリプト
        ├── test.sh            # テストスクリプト
        ├── template.sh        # テンプレート生成スクリプト
        ├── move_resolved_file.sh           # ファイル移動用スクリプト
        ├── move_resolved_file_contest.sh   # コンテスト用ファイル移動
        ├── move_resolved_file_abc.sh       # ABC用ファイル移動
        └── move_resolved_file_arc.sh       # ARC用ファイル移動
    └── template/              # テンプレートファイル
        └── main.go            # Go言語テンプレート
```

## クイックスタート

```bash
# テンプレートから新規ファイルを生成
make tml

# プログラムを実行
make run

# ビルド後に実行
make build

# テストを実行
make ts
```

## 注意事項

- 各スクリプトは `./src/script/` ディレクトリに配置されています
- `make` コマンドは Makefile のあるディレクトリから実行してください
- ファイル移動コマンドは、対象のファイルやディレクトリが存在することを前提としています
