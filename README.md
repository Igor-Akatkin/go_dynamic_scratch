# ENG
## Description:
Running dynamically built Go projects in "empty images" based on Scratch.  
In order to get a dynamically built binary file, you need to add a flag to the go build command *--ldflags "-linkmode=external"*.  
For example, ```go build --ldflags "-linkmode=external" main.go``` <br/> <br/>

If you are using a dynamically built application, you will need to add the C system library (libc), and other necessary dependencies to the base Scratch image. <br/>

Let's take the necessary utilities and libraries from the Golang image. In the context of the current example, the libraries are built from the Distroless image (gcr.io/distroless/base), as it is more lightweight, has fewer dependencies, and is originally positioned as an image for running dynamically compiled Go projects (https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md). <br/>
<br/>

## Steps to perform:
1. Build the *temp_distoless_image* image using Dockerfile from the **1_get_libs** directory. <br/>
```docker build -t temp_distroless_image``` . <br/>
This will allow us to get the necessary utilities (bash, sh, tar, gzip, find, ls, cp) from the Golang image, and add them to the distroless-image to further extract libraries from the latter.
2. Start and login to a container based on the image that was obtained using the Dockerfile described in step 1 <br/>
```docker run -it temp_distroless_image /bin/bash```
3. Execute the command to search and archive the required dependencies. <br/>
This example searches for all files with '*.so*' in their names and retrieves an archive with these files: <br/>
```find / -iname '*.so*' -print0 | tar -czvf libs.tar.gz --null -T -``` <br/>
This artefact will be required to troubleshoot dependency issues when running a dynamically built application in a Scratch-based image.
4. Take the resulting artefact out of the container, placing it on the host machine <br/>
```docker cp <CONTAINER_ID>:/libs.tar.gz <HOST_PATH>/libs.tar.gz```
5. Place the archive in the directory with the Dockerfile for building the application. <br/>
6. Using a multi-stage Dockerfile, build the application and run it in a production Scratch image to which we add the previously obtained artefact: <br/>
```docker build -t scratch-app .``` <br/>
In this example, the golang docker-image is used for the build. <br/>
There are 2 examples for two different test applications located in subdirectories of the **2_sample_go_apps** directory. Detailed information about each of them is described in the readme of the corresponding project.

___
# RUS
## Описание:
Запуск динамически собранных Go-проектов в "пустых образах", основанных на Scratch.  
Для того чтобы получить динамически собранный бинарный файл, в команду go build требуется добавить флаг *--ldflags "-linkmode=external"*.  
Например, ```go build --ldflags "-linkmode=external" main.go``` <br/> <br/>

В случае использования динамически собранного приложения, потребуется добавить в базовый Scratch-образ системную библиотеку C (libc), и другие необходимые зависимости. <br/>
Для этого из образа Golang возьмём необходимые утилиты и библиотеки. В рамках текущего примера, библиотеки собираются из Distroless-образа (gcr.io/distroless/base), так как является более легковесным, имеет меньшее количество зависимостей, и изначально позиционируется как образ под цели запуска динамически собранных Go-проектов (https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md). <br/>
<br/>
## Действия к выполнению:
1. Собрать образ *temp_distoless_image* с помощью Dockerfile из директории **1_get_libs**. <br/>
```docker build -t temp_distroless_image``` . <br/>
Это нам позволит получить из образа Golang необходимые утилиты (bash, sh, tar, gzip, find, ls, cp), и добавить их в distroless-образ для дальнейшего извлечения библиотек из последнего.
2. Запустить и войти в контейнер, основанный на образе, который был получен с помощью Dockerfile описанного в п.1 <br/>
```docker run -it temp_distroless_image /bin/bash```
3. Выполнить команду для поиска и архивирования требуемых зависимостей. <br/>
В данном примере, осуществляется поиск всех файлов, в имени которых имеется вхождение '*.so*', и получение архива с этими файлами: <br/>
```find / -iname '*.so*' -print0 | tar -czvf libs.tar.gz --null -T -``` <br/>
Данный артефакт потребуется для устранения проблем с зависимостями при запуске динамически собранного приложения в образе, основанном на Scratch.
4. Достать из контейнера полученный артефакт, разместив его на хост-машине <br/>
```docker cp <CONTAINER_ID>:/libs.tar.gz <HOST_PATH>/libs.tar.gz```
5. Архив разместить в каталоге с Dockerfile для сборки приложения. <br/>
6. С помощью multi-stage Dockerfile собрать приложение и запустить в production Scratch-образе, в который мы добавляем полученный ранее артефакт: <br/>
```docker build -t scratch-app .``` <br/>
В данном примере, для сборки используется образ golang. <br/>
Приводится 2 примера для двух различных тестовых приложений, расположенных в подкаталогах директории **2_sample_go_apps**. Подробная информация о каждом из которых, описана в readme соответствующего проекта.
