const fs = require("fs");
const path = require("path");
const readline = require("readline");
const jsonFile = "data.json";
const resultSQLFile = "dataSQL.sql";
const dataLineFile = "dataLine.json";
generateJson(dataLineFile, jsonFile);

function generateSQL(handleJsonFile) {
    fs.readFile(path.join(__dirname, handleJsonFile), "utf-8", function (
        err,
        data
    ) {
        if (err) throw err;
        let jsonData = JSON.parse(data);
        clearBeforeData(resultSQLFile);
        jsonData.forEach(function (item, index) {
            let productTaxPrice = parseFloat(item.productTaxPrice)
            productTaxPrice = productTaxPrice.toFixed(2)
            let bookSQL = `
        insert into book(id, book_name, price) values (${item.id}, '${item.book}', ${item.price});`;
            // 同步地将数据追加到文件，如果文件尚不存在则创建该文件
            fs.appendFileSync(path.join(__dirname, resultSQLFile), bookSQL, "utf-8");
        });
    });
}

function clearBeforeData(writeFileName) {
    try {
        // 检查文件是否可以读写
        // 如果可访问性检查失败，则抛出 Error
        fs.accessSync(
            path.join(__dirname, writeFileName),
            fs.constants.R_OK | fs.constants.W_OK
        );
        // 删除文件
        fs.unlinkSync(path.join(__dirname, writeFileName));
    } catch (e) {
        console.log("创建文件" + writeFileName);
    }
}

function generateJson(dataLineFile, jsonFile) {
    clearBeforeData(jsonFile);
    fs.appendFileSync(path.join(__dirname, jsonFile), "[", "utf-8");
    const rl = readline.createInterface({
        input: fs.createReadStream(path.join(__dirname, dataLineFile)),
    });

    rl.on("line", (line) => {
        lineData = line.split(/\t/g);
        if (lineData.length == 5) {
            if ("test" == lineData[0]) {
                return;
            }
            let data = {
                name: lineData[0],
                age: lineData[1],
            };
            fs.appendFileSync(
                path.join(__dirname, jsonFile),
                JSON.stringify(data) + ",\n",
                "utf-8"
            );
        }
    });

    rl.on("close", () => {
        let data = fs.readFileSync(path.join(__dirname, jsonFile), "utf8");
        data = data.slice(0, data.length - 2);
        fs.writeFileSync(path.join(__dirname, jsonFile), data);
        fs.appendFileSync(path.join(__dirname, jsonFile), "]", "utf-8");
        console.log("数据转换完成");
        generateSQL(jsonFile);
    });
}