Простое тестовое приложение, которое выводит текст "Hello" в вечном цикле с таймаутом в 1 секунду.
Также в директории есть манифест deployment-hello.yaml для создания deployment'а с нашим приложением.

P.S.
В deployment'е в качестве репозитория указан localhost:32000.
Это необходимо для использования локальных образов в microk8s (https://stackoverflow.com/questions/55297278/how-to-use-local-docker-images-with-microk8s):
1) microk8s имеет частный реестр, который можно использовать для этой цели.
Необходимо включить реестр заранее, с помощью команды:
$ microk8s.enable registry

2) Реестр сопоставляет трафик с портом 32000, поэтому вам придется отправить образ докера в реестр. 
Если образ уже присутствует в локальном каталоге, вы можете использовать команду docker tag:
$ docker tag <imageName:version> localhost:32000/<imageName:version>

Например: 
$ docker tag hello-app-image:latest localhost:32000/hello-app-image:latest

3) Затем выполнить push:
$ docker push localhost:32000/<imageName:version>.

Например:
$ docker push localhost:32000/hello-app-image:latest

Подробнее в оф. документации https://microk8s.io/docs/registry-built-in