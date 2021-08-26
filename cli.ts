#!/usr/bin/env node


import fs from "fs";

interface PostItem {
  path: string;
  date: Date;
  title: string;
  description: string;
  url: string;
}


const packagejson = JSON.parse(fs.readFileSync("./package.json", "utf8"));

