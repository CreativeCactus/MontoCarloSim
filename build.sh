docker build -t monte .
docker run --name monteCarlo --rm -itv "$PWD":/usr/src/carlo -w /usr/src/carlo -p 3000:3000 monte bash e2e.sh "$@" # buildonly testonly all
