if [ ! -f "output/logs" ]; then 
    mkdir -p "output/logs"
fi

cp -r conf output/ 

conf_dir=output/conf log_dir=output/log exec output/bin/search_codimd 