const {
  Sequelize,
  Model,
  DataTypes
} = require('sequelize')

var sequelizeInstance = require('../config/db');

class Contents extends Model {

}

Contents.init({
  id: {
      type: Sequelize.BIGINT(20),
      primaryKey: true,
      allowNull: false,
      unique: true,
      autoIncrement: true
  },
  chapter_id: Sequelize.BIGINT(20),
  content: Sequelize.TEXT,
  created_at: {
    type: Sequelize.DATE,
    defaultValue: Sequelize.NOW,
    name: 'createdAt',
    field: 'created_at'
  },
  updated_at: {
    type: Sequelize.DATE,
    defaultValue: Sequelize.NOW,
    name: 'updatedAt',
    field: 'updated_at'
  },
  deleted_at: {
    type: Sequelize.DATE,
    name: 'deletedAt',
    field: 'deleted_at'
  },
},{
  sequelize: sequelizeInstance,
  tableName: 'contents',
  timestamps: false
})

module.exports = { Contents }
