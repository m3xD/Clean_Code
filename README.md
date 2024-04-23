# Clean Code

## SOLID
### Single responsibility principle:
- Các class chỉ nắm giữ một chức năng, chỉ sửa class với một mục đích duy nhất.
### Open-Closed principle
- Có thể mở rộng class thông qua các interface, nhưng không sửa lại class đã hoạt động tốt.
### Liskov substitution principle:
- Trong một chương trình, các object của class con có thể thay thế class cha mà không làm thay đổi bản chất của class cha.

## Stagery Pattern:
### 1. Đặt vấn đề:
- Khi class muốn thêm các methods vào một class được cài đặt trước.
- Các methods khi được thêm vào thì lại nảy sinh vấn đề: các subclass kế thừa class đó lại có các hành vi không mong muốn.
- Các methods không cần thiết thì bị lặp lại ở các subclass không cần methods đấy.
- Sử dụng interface cũng chưa phải một phương án tốt vì các subclass trước đấy phải implement.
### 2. Giải pháp:
- Ta li khai các methods cần thay đổi ra một class khác, tránh thay đổi code được implement sẵn trong class hiện thời.
  ```
  Ví dụ: Con vịt thì có thể bay và có thể kêu.
  --> ta có 2 class là Fly và Quack
  ```
- Sử dụng interface để thể hiện một tập behavior, từng class sẽ implement interface này để thể hiện rõ chức năng cụ thể.
  ```
  Ví dụ: Ta sử dụng 2 interface là Flyable và Quackable
  --> Các class có thể implement 2 interface là : Fly, NoneOfFly, Quack, NoQuackAtAll
  ```
- Thêm các class vừa li khai theo quan hệ HAS-A vào trong class cần thêm methods.
- Ở các subclass thì các thuộc tính vừa được thêm sẽ được khởi tạo theo hành vi phù hợp như Fly hoặc NoneOfFly,...
### 3. Điểm mạnh:
- Class hiện thời không cần biết các hành vi được thêm thực hiện như nào, mọi thứ được khởi tạo dưới các subclass.
- Các chức năng được thêm vào mà có sự thay đổi độc lập với code mà không có sự thay đổi, giảm bớt bug, tăng khả năng bảo trì.

## Observer Pattern
### 1. Đặt vấn đề:
- Có nhiều object cần nhận biết sự thay đổi của một object cụ thể nào đó.
- Số lượng các object cần notify có sự thay đổi liên tục.
### 2. Giải pháp:
- Dựa vào quan hệ 1-nhiều, các dependents sẽ nhận được notification nếu có sự thay đổi từ object.
- Các depedents có thể lựa chọn subcribe hay unsubcribe object.
![pic_of_observer_patter](https://learning.oreilly.com/api/v2/epubs/urn:orm:book:9781492077992/files/assets/f0052-01.png)
- Sẽ có 2 interface đại diện cho subject và observer.
- Interface subject sẽ gồm các methods reg, rem, notify giúp thêm, xoá, thông báo cho các observers.
- Interface observer có method update giúp cập nhật trạng thái được push từ subject.
### 3. Key:
- Giữa object và observer thường không biết nhiều thông tin về nhau, tính mềm dẻo được thể hiện.
- Subject chỉ biết về các observers là các observers implement các interface observer.
- Vì thế mình có thể thêm các observers mà subject không cần biết gì về chúng.
- Chả cần phải thay đổi gì bên trong object để có thể thêm observer.
- Việc tạo ra các concreate class objec và observer khá tự do, có thể tái sự dụng được.
- Vì chúng có liên kết không chặt chẽ, nên việc thay đổi bên object hay observer không làm thay đổi bên đối diện.
>[!NOTE]
>Strive for loosely coupled designs between objects that interact.
>Điều này giúp build system mềm dẻo, có thể tạo ra các thay đổi mà gây ra một tác động không đáng kể đến các implement hiện thời.
### 4. Bonus:
- Các observers có thể chủ động pull dữ liệu, chọn lọc các dữ liệu cần thay vì lấy một cục dữ liệu dư thừa.
- Bằng cách implement các methods get trong object, observers giờ đây tự do hơn.

## Decorator Pattern:
### 1. Đặt vấn đề:
- Khi muốn thêm các tính năng mới cho một object mà không làm thay đổi các object này.
- Khi muốn việc thay đổi tính năng thực hiện trong quá trình run-time.
- Khi các đối tượng kh thể thêm tính năng bằng mối quan hệ kế thừa.
### 2. Giải pháp:
- Sử dụng một interface quy định các methods chung cho tất cả các component tham gia.
- Một abstract class sẽ implement interface gọi là Decorator, các decorators sẽ có một trường để lưu trữ giá trị đối tượng gốc. Trong các Decorators, ta có thể bổ sung thêm các tính năng cần thiết.
- Các concreate class được extends/implement từ Decorators sẽ implement các tính năng mới từ Decorators.
![decorators](https://learning.oreilly.com/api/v2/epubs/urn:orm:book:9781492077992/files/assets/f0101-01.png)
- Ví dụ: ta có thể đọc file txt bằng cách sử dụng concreate decorator là FileInputStream, đồng thời muốn tận dụng các tính năng khác, ví dụ: BufferedInputStream.
```
  InputStream in = new FileInputStream('smt.txt');
  in = new BufferedInputStream();
```
### 3. Điểm mạnh:
- Có khả năng thêm các tính năng trong quá trình run-time.
- Đối tượng có thể wrapper nhiều lần.

## Factory Abstract Pattern:
### 1. Đặt vấn đề:
- Ta muốn khởi tạo các đối tượng mà không để lộ phần logic khởi tạo đối tượng với client.
- Được sử dụng khi ta có một lớp cha với nhiều lớp con, cần phải trả về object dựa vào đầu vào.
### 2. Giải pháp:
- Bổ sung thêm một lớp Factory có nhiệm vụ sử dụng if-else hoặc switch-case để xác định class đầu ra.
- Super class có thể là class, interface hoặc là abstract class.
- Subclass sẽ implement các methods của supperclass. 
![factory](https://learning.oreilly.com/api/v2/epubs/urn:orm:book:9781492077992/files/assets/f0156-01.png)
- Ta có thể tạo thêm các concreate factory class nếu có nhiều superclass.
### 3. Điểm mạnh:
- Giảm sự phụ thuộc giữa các class, code ở phía client sẽ không thay đổi cho dù có sự thay đổi ở subclass hay factory class.
- Nới rộng code hơn khi có nhiều superclass bằng cách extends/implement factory class.
- Khởi tạo các object và che dấu logic.

### Singleton Pattern:
### 1. Đặt vấn đề:
- Mong muốn có một đối tượng duy nhất để có thể truy xuất mọi nơi.
- Sử dụng global variable nhưng vi phạm qui tắc tính đóng gói của OOP.
### 2. Giải pháp:
- Sử dụng class có private constructor để chỉ có thể khởi tạo trong class.
- Đặt biến có phạm vi private.
- Sử dụng getter để có thể lấy được instance.
### 3. Điểm mạnh:
- Chỉ có 1 instance duy nhất để truy cập.
- Có thể sử dụng cho các desgin pattern khác như Factory Abstract.

### Proxy Pattern:
### 1. Đặt vấn đề:
- Mong muốn bổ sung một số thao tác khi thực hiện trên class thực (xác thực,..)
- Khi muốn thực hiện remote trên môi trường mạng.
- Cần nâng cao hiệu năng tải.
### 2. Giải pháp:
- Proxy đứng ra làm đại diện cho class thực để đảm bảo một số yêu cầu hay điều kiện.
- Tất cả các truy cập đến class thực đều phải thông qua proxy.
- Proxy định nghĩa lại một số hành vi sao cho phù hợp với mong muốn người sử dụng, đồng thời thực thi những methods trên class thực.
- Đồng nghĩa proxy nắm giữ references của class thực (quan hệ HAS-A).
- Có nhiều phân loại proxy như: vitural proxy, protection proxy, remote proxy,..., các proxy này sẽ phục vụ mục đích cụ thể của người sử dụng.
- Nhìn chung thì các loại proxy đều có đặc điểm chung như:
  * Cung cấp khả năng truy cập gián tiếp qua proxy.
  * Reference đến đối tượng đích.
  * Cả reference và proxy đều có cùng một interface.
### 3. Điểm mạnh:
- Có thể thiết lập nhiều thao tác trước khi thực hiện trên class đích.
- Giảm chi phí khi chỉ thực sự khởi tạo class đích.
- Tăng khả năng chịu tải.
- Dễ nâng cấp, bảo trì.

