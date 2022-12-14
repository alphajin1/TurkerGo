# 4장 변수

## 변수의 4가지 속성
```bash
var a = 10
```
1. 이름 : 개발자를 위한 것 (a)
2. 값 : 변수가 들고 있는 것 (10)
3. 주소 : 컴퓨터를 위한 것, 시작 주소
4. 타입 : 
   1. 타입별로 사이즈가 다름
   2. 시작 주소와 사이즈를 어느 공간에 값이 있는지 알 수 있음

### 숫자 타입
- uint8, unit16, uint32, uint64
- int8, ... int64
- float32, float64
---
- byte (uinit8과 같음, aliasing)
- rune (int32와 같음, aliasing)
  - 문자 1개(UTF-8, 1~3 byte)를 나타냄
  - rune 이 모이면 string이 됨
- int (32bit computer에서는 int32, 64bit는 int64)
- uinit (int와 같은 방식)

### 그외 타입
- bool
- string
- 배열
- 슬라이스
- 구조체
- 포인터
- 함수타입
- 맵
- 인터페이스

### 별칭 타입
- rune = int32, byte = uint8

### 여러가지 변수 선언법
```bash
# 초기값 O
var a int = 10

# 초기값 X, default = 0
var a int

# type 생략, 초기값이 반드시 있어야, 숫자는 기본이 int
var a = 10

# 선언 대입문 (:=), 요렇게 많이 쓴다고 함.
a := 10

# type 생략, float64가 기본타입
f := 3.14 
```

### 타입별 기본값
pass

### 타입 변환
연산의 각 항목의 타입은 **반드시** 같아야 한다.
