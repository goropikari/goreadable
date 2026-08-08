# goreadable

Go コードから、AI または人間が可読性をレビューすべき候補を抽出する CLI です。
可読性を単一スコアで合否判定するのではなく、計測値・閾値・理由・ソース断片を提示します。

## 使い方

```sh
# カレントパッケージを解析（既定はテキスト出力）
go run ./cmd/goreadable .

# リポジトリ全体を JSON で解析
go run ./cmd/goreadable --format json ./...

# 閾値を指定して Git 差分を解析
go run ./cmd/goreadable --diff HEAD --max-function-lines 40 ./...
```

検出候補があっても終了コードは `0` です。候補はレビューの優先順位付けに使う情報であり、自動的な不合格判定ではありません。オプション不正、設定ファイル不正、解析エラーは非 `0` で終了します。

## 検出ルール

関数について、次の値を計測します。

- 行数（`--max-function-lines`、既定値 80）
- ネスト深度（`--max-nesting-depth`、既定値 4）
- 循環的複雑度（`--max-cyclomatic-complexity`、既定値 10）
- 引数数（`--max-function-args`、既定値 5）

構造体・型について、次の値を計測します。

- 構造体フィールド数（`--max-struct-fields`、既定値 8）
- 型に関連するメソッド数（`--max-type-methods`、既定値 10）

`*_test.go` も解析対象で、結果には本番コード（`production`）またはテストコード（`test`）の区分が含まれます。生成コードと `vendor/` は既定で除外します。

## 設定

解析対象のルートに `goreadable.json` を置くと、閾値を設定できます。CLI フラグ、設定ファイル、既定値の順に優先されます。

```json
{
  "thresholds": {
    "function_lines": 60,
    "nesting_depth": 3,
    "cyclomatic_complexity": 8,
    "function_arguments": 4,
    "struct_fields": 6,
    "type_methods": 8
  }
}
```

## 出力

`--format text`（既定）は、候補名・ファイル位置・コード区分・検出理由を人間向けに表示します。`--format json` は次の情報を含むバージョン付きレポートを出力します。

- `kind`、`name`、`path`、`start_line`、`end_line`
- `code_kind`（`production` または `test`）
- 計測値（`metrics`）と適用閾値（`thresholds`）
- 候補理由（`reasons`）
- AI が追加の読込なしに一次評価できる対象ソース（`source`）

JSON は後続の AI レビュー工程へ渡すための入力として利用できます。goreadable 自身は外部 AI API を呼び出しません。

## 開発

```sh
go test ./...
go test -race ./...
make fmt
make lint
```

受入ハーネスは [.acceptance-harness/manifest.json](.acceptance-harness/manifest.json)、受入テストは [tests/acceptance/goreadable_test.go](tests/acceptance/goreadable_test.go) にあります。
