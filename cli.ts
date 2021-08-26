#!/usr/bin/env node

import { Command } from "commander";

import { prompt } from "inquirer";

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

const renderBlogPostList = async (posts: PostItem[]) => {
  const answers = await prompt([{
    type: 'list',
    name: 'selected',
    loop: false,
    message: 'Which post do you want to read?',
    choices: posts.map(post => { return { name: post.title, value: post }; }),
  }]);

  await renderBlogPost(answers.selected);

  const ask = await prompt([{
    type: 'confirm',
    name: 'again',
    message: 'Want to read another blog post (just hit enter for YES)?',
    default: true,
  }]);

  if (!ask.again) {
    process.exit(0);
  }

  await renderBlogPostList(posts);
}

const renderBlogPost = async (post: PostItem) => {
  const response = await fetch(post.path)
  const md = await response.text()

  console.log(marked(md));
};

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
  .then(async (json: PostItem[]) => { await renderBlogPostList(json.reverse()); })
  .catch(error => console.log(error));
