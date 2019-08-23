var fs = require('fs');
var path = require('path');
var handleJsonFile = "data.json";
var resultSQLFile = "Book.sql";
fs.readFile(path.join(__dirname, handleJsonFile), 'utf-8', function (err, data) {
    if (err) throw err;
    var jsonData = JSON.parse(data);
    clearBeforeData();
    jsonData.forEach(function (item, index) {
        var bookSQL = `
        insert into book(id, book_name, price)
values (${item.id}, '${item.book}', ${item.price});`;
        // 同步地将数据追加到文件，如果文件尚不存在则创建该文件
        fs.appendFileSync(path.join(__dirname, resultSQLFile), bookSQL, 'utf-8');
    })
});

function clearBeforeData() {
    try {
        // 检查文件是否可以读写
        // 如果可访问性检查失败，则抛出 Error
        fs.accessSync(path.join(__dirname, resultSQLFile), fs.constants.R_OK | fs.constants.W_OK);
        // 删除文件
        fs.unlinkSync(path.join(__dirname, resultSQLFile));
    } catch (e) {
        console.log("创建文件" + resultSQLFile);
    }
}
