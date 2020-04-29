const cheerio = require('cheerio')
const request = require('superagent')
require('superagent-charset')(request)
const { config } = require('../config')

async function getBookList() {
  const res = await request.get(config.baseUrl + '/files/article/html/2/2689/index.html').charset('gbk')
  const arr = []
  let $ = cheerio.load(res.text, { decodeEntities: false })
  $('.box .zjlist4 li a').each((idx, ele) => {
    let news = {
      id: idx,
      title: $(ele).text(),
      href: $(ele).attr('href')
    };
    arr.push(news)
  })
  console.log('arr', arr)
}

module.exports = {
  getBookList
}