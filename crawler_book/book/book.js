const fs = require('fs')
const path = require('path')

const cheerio = require('cheerio')
const request = require('superagent')
require('superagent-charset')(request)
const { config } = require('../config')

const addr = path.join(__dirname, '../chapter/1.json')


async function getChapterList() {
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
  const chapter = JSON.stringify(arr, null, 4)
  // console.log('chapter: ', chapter)
  fs.writeFileSync(addr, chapter)
}

function getContent() {
  const json = fs.readFileSync(addr, 'utf-8')
  const res = JSON.parse(json)
  return res
}

async function getContentText(href, chapterId, novelId) {
  const contentAddr = path.join(__dirname, '../content/' + novelId)
  const contentPath = path.join(__dirname, '../content/' + novelId + '/' + chapterId + '.txt')
  fs.exists(contentAddr,(exists) => {
    if (!exists) {
      fs.mkdirSync(contentAddr)
    }
  })
  const res = await request.get(config.baseUrl + '/files/article/html/2/2689/' + href).charset('gbk')
  let $ = cheerio.load(res.text, { decodeEntities: false })
  $('#htmlContent').find('table').remove()
  const txt = $('#htmlContent').html()
  fs.writeFileSync(contentPath, txt)
}

module.exports = {
  getChapterList,
  getContent,
  getContentText
}