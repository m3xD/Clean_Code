# Clean Code

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
