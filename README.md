## Go と Python で解き比べ

### 問題設定
- [Python ではじめる数理最適化（第2版）](https://www.ohmsha.co.jp/book/9784274231759/) の 第２章 P 17 の問題
- Go と Python で解き比べ


### 前提
以下がインストールされていること
- uv
- Go

### 準備

```sh
$ git clone https://github.com/ysaito8015/simplex-algoritm-example
```


### Python

```sh
$ cd simplex-algoritm-example/python
$ uv run solve.py
```


```sh
Status: Optimal
x=  18.0 y=  4.0 obj=  26.0
```


### Go

```sh
$ cd simplex-algoritm-example/go
$ go mod tigy
$ go run ./main.go
```


```sh
x= 18 y= 4 obj= 26
```
