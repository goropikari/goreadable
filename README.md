# goreadable

Go コードの全関数について可読性メトリクスを出力し、必要に応じて AI または人間がレビューすべき候補へ絞り込める CLI です。
可読性を単一スコアで合否判定するのではなく、計測値・閾値・理由・ソース断片を提示します。

## 使い方

`goreadable` を利用可能な状態にしてリポジトリを clone したら、まず同梱サンプルをそのまま実行できます。

```sh
goreadable --thresholds-only --max-function-lines 5 --max-function-args 4 ./samples/readme
function BuildReport (samples/readme/main.go:3-11, production)
  - function_arguments=6 exceeds threshold 4
  - function_lines=9 exceeds threshold 5
```

この例は、引数が 6 個で 9 行の `BuildReport` を含む固定サンプルを解析します。実行環境やリポジトリ内の他のコードに依存せず、可読性レビュー候補とその理由を確認できます。

ほかの対象を解析する場合は、パスを置き換えます。

```sh
# カレントパッケージを解析（既定はテキスト出力）
goreadable .

# リポジトリ全体を JSON で解析
goreadable --format json ./...

# 閾値を超えた候補だけを Git 差分から解析
goreadable --thresholds-only --diff HEAD --max-function-lines 40 ./...

# パッケージ内の全関数のメトリックを表示（既定）
goreadable ./...

# 指定したパッケージ関数（またはメソッド）のメトリックを表示
goreadable --function analysis.Analyze --function analysis.Options.MetricsOnly ./...
```

### 出力例

AI レビューへ渡す場合は、同じ固定サンプルに JSON を指定します。

```sh
goreadable --format json --thresholds-only --max-function-lines 5 --max-function-args 4 ./samples/readme > readability-candidates.json
```

出力された `readability-candidates.json` には、候補のソース断片・計測値・閾値・理由が含まれます。

検出候補があっても終了コードは `0` です。候補はレビューの優先順位付けに使う情報であり、自動的な不合格判定ではありません。オプション不正、設定ファイル不正、解析エラーは非 `0` で終了します。

## 検出ルール

関数について、次の値を計測します。

- 行数（`--max-function-lines`、既定値 80）
- ネスト深度（`--max-nesting-depth`、既定値 4）
- 循環的複雑度（`--max-cyclomatic-complexity`、既定値 10）
- 引数数（`--max-function-args`、既定値 5）
- ローカル変数数（`--max-local-variables`、既定値 15）
- 制御ブロック数（`--max-control-blocks`、既定値 8）
- return 文の数（`--max-return-points`、既定値 5）
- 論理演算子数（`--max-boolean-operators`、既定値 8）
- 1 条件式の最大項数（`--max-condition-terms`、既定値 4）
- 関数・メソッド呼び出し数（`--max-function-calls`、既定値 15）
- 意味のあるリテラル数（`--max-literal-values`、既定値 10）
- 関数リテラル数（`--max-closures`、既定値 2）
- コメント行数（`--max-comment-lines`、既定値 10）
- 文の数（`--max-statements`、既定値 40）
- シグネチャの型依存数（`--max-type-dependencies`、既定値 5）

ローカル変数数は、関数内の `var`、`:=`、`range` と型 switch の短縮宣言で導入される名前を数えます。引数と結果値は含めません。制御ブロック数は `if`、`for`、`range`、`switch`、型 switch、`select` の数です。条件項数は `&&` と `||` で分けられた項の最大数です。リテラル数は `0`、`1`、空文字列を除外し、型依存数はレシーバー・引数・戻り値に現れる異なる型式の数です。関数本文内のクロージャの内部は、本文の各メトリクスから除外します。

構造体・型について、次の値を計測します。

- 構造体フィールド数（`--max-struct-fields`、既定値 8）
- 型に関連するメソッド数（`--max-type-methods`、既定値 10）
- 公開メンバー数（`--max-exported-members`、既定値 10）

公開メンバー数は、構造体の公開フィールドと型の公開メソッドの合計です。

`*_test.go` も解析対象で、結果には本番コード（`production`）またはテストコード（`test`）の区分が含まれます。生成コードと `vendor/` は既定で除外します。

宣言の直前に `// goreadable:ignore` または `/* goreadable:ignore */` を置くと、その関数・型を候補から除外できます。除外コメントは、意図的に複雑なコードなど、静的な閾値ではレビュー対象にしたくない宣言に使用してください。

## 設定

解析対象のルートに `goreadable.json` を置くと、閾値を設定できます。CLI フラグ、設定ファイル、既定値の順に優先されます。

```json
{
  "thresholds": {
    "function_lines": 60,
    "nesting_depth": 3,
    "cyclomatic_complexity": 8,
    "function_arguments": 4,
    "local_variables": 12,
    "control_blocks": 6,
    "return_points": 4,
    "boolean_operators": 6,
    "max_condition_terms": 3,
    "function_calls": 12,
    "literal_values": 8,
    "closure_count": 2,
    "comment_lines": 8,
    "statement_count": 30,
    "type_dependencies": 4,
    "struct_fields": 6,
    "type_methods": 8,
    "exported_members": 8
  }
}
```

## 出力

`--format text`（既定）は、関数名・ファイル位置・コード区分・全メトリックを人間向けに表示します。`--format json` は次の情報を含むバージョン付きレポートを出力します。

- `kind`、`name`、`path`、`start_line`、`end_line`
- `code_kind`（`production` または `test`）
- 計測値（`metrics`）と適用閾値（`thresholds`）
- 候補理由（`reasons`）

JSON は後続の AI レビュー工程へ渡すための軽量な入力として利用できます。goreadable 自身は外部 AI API を呼び出しません。

既定では、閾値を超えているかどうかにかかわらず、対象パッケージ内の全関数を出力します。`--thresholds-only` を指定すると、従来どおり閾値を超えた関数・型だけをレビュー候補として出力します。`--function` は繰り返し指定でき、`パッケージ名.関数名`（メソッドは `パッケージ名.型名.メソッド名`）で特定の関数だけを出力します。全関数出力では、テキスト出力にも各関数のメトリックを表示します。

## 開発

```sh
go test ./...
go test -race ./...
make fmt
make lint
```

受入ハーネスは [.acceptance-harness/manifest.json](.acceptance-harness/manifest.json)、受入テストは [tests/acceptance/goreadable_test.go](tests/acceptance/goreadable_test.go) にあります。
