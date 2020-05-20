const axios = require("axios");
const cheerio = require("cheerio");

// Be careful about loading the entire dataset, it takes forever!
// const url = "https://www.aonprd.com/Monsters.aspx?Letter=All";
const url = "https://www.aonprd.com/Monsters.aspx?Letter=Z";

const fetchData = async () => {
  const res = await axios.get(url);
  return cheerio.load(res.data);
};

const run = async () => {
  const $ = await fetchData();
  console.log($("#main-wrapper").html());
};

run();
