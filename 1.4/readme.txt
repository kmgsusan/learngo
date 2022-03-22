# naked return 
    해당 return 은 return 명시를 안하고 그냥 'return' 만하면 되는기능으로 아래처럼 작성하면됨. 변수명이랑 자료형 명시하면 된다.
    func leng(name string) (len int, upper string)

# defer
    함수내에서 쓰는것으로 함수내 코드가 모두 실행이 완료된이후 함수가 종료되는 시점에 실행되는 코드를 명시하면 된다.