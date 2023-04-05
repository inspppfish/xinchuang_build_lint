cat says | xargs -I {} -P 4 sh -c 'cowsay {} > {}'
