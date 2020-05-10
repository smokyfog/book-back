const express = require('express')
const app = express()
const { getChapterList, getContent, getContentText } = require('./book/book') 
const { Chapters } = require('./db/schema/chapter')

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
    getContentText(item.href, item.id, 1)
    console.log('爬取' + item.id + '结束')
  }
  res.send({
    code: 0,
    msg: json[0]
  })
})

app.get('/save/chapter', async (req, res) => {
  const json = await getContent()
  json.map(item => {
    const chapter = {
      id: item.id,
      name: item.title,
      novel_id: 1,
      chapter_num: 1,
      content_id: 1
    }
    Chapters.create(chapter)
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