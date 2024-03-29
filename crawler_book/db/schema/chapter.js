const {
  Sequelize,
  Model,
  DataTypes
} = require('sequelize')

var sequelizeInstance = require('../config/db');

class Chapters extends Model {
}

Chapters.init({
  id: {
      type: Sequelize.BIGINT(20),
      primaryKey: true,
      allowNull: false,
      unique: true,
      autoIncrement: true
  },
  novel_id: Sequelize.BIGINT(20),
  name: Sequelize.STRING(255),
  chapter_num: Sequelize.BIGINT(20),
  content_id: Sequelize.BIGINT(20),
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
  tableName: 'chapters',
  timestamps: false
})

module.exports = { Chapters }
