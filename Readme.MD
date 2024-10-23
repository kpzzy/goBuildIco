## go build 시 exe 파일에 icon 을 추가하는 방법

## 문제
- build 된 exe 파일에 아이콘을 추가해달라는 요청
- 우클릭 속성에서 아이콘 변경이 가능한걸로 기억되었는데 어째서인지 되지 않음

## 해결
build 시 icon을 추가하는 방법은 다음과 같다.

.ico 파일을 go 폴더에 넣은 후 rsrc package 를 통해 syso 생성, 이후 build 하면 icon 이 추가되어 생성된다.

## 실습

오늘 테스트 할 이미지는 대한민국 국기모양의 파일이다.

![image](https://github.com/user-attachments/assets/1cf09436-537f-4b81-a5a8-c4252452132c)


먼저 기본적으로 go build 시 아래와같이 아이콘이 없는 exe 파일이 생긴다

```
// 먼저 package 를 설치한다
go install github.com/akavel/rsrc@latest

// 이후 rsrc 명령어로 syso 파일을 생성한다
rsrc -ico favicon.ico -o rsrc.syso
```

현재 폴더구조는 다음과 같다

![image](https://github.com/user-attachments/assets/0a27b2e3-a4e1-4111-bda6-402b2c9c557d)

이제 go build를 한다

```
go build
```

![image](https://github.com/user-attachments/assets/9a8e53c1-2cd5-452a-b46f-6a50706b2d2f)

favicon.ico 의 이미지로 exe 파일이 생성된다

## 주의점

go build 시 -o 명령어로 파일명을 입력하면 syso 인식하지못하기 때문에 go build 가 적합하다

ex) 실패 command = go build test.exe -o main.go
ex) 성공 command = go build