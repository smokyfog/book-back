const fs = require('fs')
const path = require('path')

const express = require('express')
const app = express()
const { getChapterList, getContent, getContentText } = require('./book/book') 
const { Chapters } = require('./db/schema/chapter')
const { Contents } = require('./db/schema/content')

app.get('/book', async (req, res) => {
  getChapterList()
  res.send({
    code: 0
  })
})

app.get('/read_book', async (req, res) => {
  const json = await getContent()
  // await getContentText(json[0].href, json[0].id, 1)
  for (const item of json) {
    console.log('正在爬取', item.id)
    getContentText(item.href, item.id, 2)
    console.log('爬取' + item.id + '结束')
  }
  res.send({
    code: 0,
    msg: json[0]
  })
})

app.get('/save/chapter', async (req, res) => {
  const json = await getContent()
  json.map(async (item, idx) => {
    const chapter = {
      name: item.title,
      novel_id: 1,
      chapter_num: idx + 1,
      content_id: 1
    }
    try {
      const chaRes = await Chapters.create(chapter)
      const contentPath = path.join(__dirname, './content/1/' + item.id + '.txt')
      const con = fs.readFileSync(contentPath, 'utf-8')
      const content = {
        chapter_id: chaRes.dataValues.id,
        content: con.trim()
      }
      const ConRes = await Contents.create(content)
      const UpdateRes = await Chapters.update({content_id: ConRes.dataValues.id}, {where: {id: chaRes.dataValues.id}})
      console.log('UpdateRes', UpdateRes)
      console.log('id为 ' + item.name + ' 的小说已存入， contentId为' + ConRes.dataValues.id,)
    } catch (err) {
      console.log('save chapter error : ', err, chapter)
    }
  })
  res.send({
    query: req.query
  })
})

const server = app.listen(8899, () => {
  const host = server.address().address
  const port = server.address().port
  console.log('Your App is running at http://%s:%s', host, port)
})