version=`git symbolic-ref --short -q HEAD`

image=`basename $PWD`:$version

aliyunImage=registry.cn-hangzhou.aliyuncs.com/everywan/$image

docker rmi --force aliyunImage 2> /dev/null

docker build -t $aliyunImage .

docker push $aliyunImage
