git config --global https.https://github.com.proxy https://127.0.0.1:41091
git config --global http.https://github.com.proxy https://127.0.0.1:41091
git push
git config --global --unset https.https://github.com.proxy
git config --global --unset http.https://github.com.proxy