### Khởi tạo biến
syntax: < kiểu dữ liệu > < tên biến > {[giá trị khởi tạo]}
``` flux
text name {"John"}
number age {25} 
boolean isStudent {true}
ipv4 localHost {127.0.0.1}
```
### Lưu ý: 
    Giá trị khởi tạo phải trùng khớp với kiểu dữ liệu của biến
    Có thể không cần giá trị khởi tạo

### Khởi tạo hàm
syntax:  < tên hàm > {
([parameter <tên tham số>] )

<Nội dung hàm>

}
``` flux
myFun {

(parameter isMember)

return isMember ? '$2.00' : '$10.00';

}

myFun(isMember(true)) //$2.00

myFun(isMember(false))//$10.00
```