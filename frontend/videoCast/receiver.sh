cd /usr/share/nginx/html && npm i && yes | npx webpack
rm /usr/share/nginx/html/*momo*

/usr/sbin/nginx -g "daemon off;"
