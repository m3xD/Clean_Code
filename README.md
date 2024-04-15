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
