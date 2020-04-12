# client/proto/*.jsの先頭行に /* eslint-disable */ を追加
# Macはgsed, Linuxの場合はsedに変更
find client/src/generated/*.js | xargs gsed -i '1i /* eslint-disable */'