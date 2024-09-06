FROM scratch

# goreleaserが作成したバイナリをコピー
COPY dbreseter-by-echo /

# バイナリを実行するためのエントリーポイント
ENTRYPOINT ["/dbreseter-by-echo"]

# アプリケーションがリッスンするポートを指定
EXPOSE 8080
