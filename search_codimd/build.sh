
if [ ! -f output/bin ]; then 
    mkdir -p output/bin
fi

go build  -o output/bin/search_codimd main.go


