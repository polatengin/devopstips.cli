#!/usr/bin/env node

import { Command } from "commander";


import fs from "fs";

interface PostItem {
  path: string;
  date: Date;
  title: string;
  description: string;
  url: string;
}


const packagejson = JSON.parse(fs.readFileSync("./package.json", "utf8"));

const program = new Command();

program
  .version(packagejson.version)
  .description(packagejson.description)

