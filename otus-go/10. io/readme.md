### Последовательный и произвольный доступ

Устройства/технологии ввода/вывода данных можно условно разделить на поддерживающие произвольный доступ
* жесткие диски
* память


и поддерживающие последовательный доступ
* терминал
* сетевое соединение
* pipe


Как следствие есть два набора интерфейсов
* io.Reader , io.Writer - для последовательного доступа
* io.ReaderAt , io.WriterAt , io.Seeker - для произвольного доступа