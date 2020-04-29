const express = require('express')
const app = express()
const { getBookList } = require('./book/book') 

app.get('/book', async (req, res) => {
  getBookList()
  res.send({
    code: 0
  })
})

const server = app.listen(8899, () => {
  const host = server.address().address
  const port = server.address().port
  console.log('Your App is running at http://%s:%s', host, port)
})