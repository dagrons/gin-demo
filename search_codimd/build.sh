if [ ! -f "output/bin" ]; then 
    mkdir -p "output/bin"
fi

go build  -o output/bin/ app.go