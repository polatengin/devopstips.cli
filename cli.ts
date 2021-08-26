#!/usr/bin/env node

import { Command } from "commander";


import marked from 'marked';
import TerminalRenderer from 'marked-terminal';

import fs from "fs";
import fetch from "node-fetch";

interface PostItem {
  path: string;
  date: Date;
  title: string;
  description: string;
  url: string;
}

marked.setOptions({
  renderer: new TerminalRenderer()
});


const packagejson = JSON.parse(fs.readFileSync("./package.json", "utf8"));

const program = new Command();

program
  .version(packagejson.version)
  .description(packagejson.description)

program
  .command("read [url] [title]")
  .alias("r")
  .description("Read a blog post")
  .action((url, title) => {
    console.log("Read " + url + " as " + title);
  });

program
  .command("random")
  .alias("x")
  .description("Read a random blog post")
  .action(() => {
    console.log("Reading a random blog post");
  });

fetch("https://devopstips.net/api/posts.json")
  .then(response => response.json())
  .catch(error => console.log(error));
