const Sequelize = require('sequelize');
// 数据库配置文件

var sqlConfig = {
    host: "47.93.10.229",
    user: "root",
    password: "Adsxlwh_2008",
    database: "test"
}


const sequelizeInstance = new Sequelize(sqlConfig.database, sqlConfig.user, sqlConfig.password, {
    host: sqlConfig.host,
    dialect: 'mysql',
    pool: {
        max: 10,
        min: 0,
        acquire: 30000,
        idle: 10000
    },
    dialectOptions: {
        charset: "utf8mb4",
        collate: "utf8mb4_unicode_ci",
        supportBigNumbers: true,
        bigNumberStrings: true
    },
    define: {
        underscored: true,
        charset:'utf8mb4'
    },
    freezeTableName: true,
    timezone: '+08:00'
});


// 创建模型
sequelizeInstance.sync({
    force: false
})


module.exports = sequelizeInstance