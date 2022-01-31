<p align="center">
  <a href="https://npmjs.com/package/retonio" target="_blank" rel="noopener noreferrer">
    <img width="230" src="https://ik.imagekit.io/vrfrbvdn0j/sddev/GoScission.svg?ik-sdk-version=javascript-1.4.3&updatedAt=1643638979487" alt="Retonio logo">
  </a>
</p>
<br/>
<p align="center">
  <!-- <a href="https://npmjs.com/package/retonio"><img src="https://badgen.net/npm/v/retonio" alt="npm package"></a> -->
  <a href="https://github.com/sdoerger/GoScission"><img src="https://badgen.net/github/release/sdoerger/GoScission/" alt="github project"></a>
</p>
<br/>

# GoScission
A simple script to split files into an unknown folder amount, with a limit per folder.

<!-- - 📦 Wrapps a Pinia Store in one line
- 🔧 Customizeable
- ⏳ Loading States -->

## Usage

Place the bin file GoScission whereever you like.
Run it and add as parameters the path of the folder where your files are and the folder limit, i. e.:

`./GoScission /home/myDirectore 5`

## Result

It creates a subfolder (gosci) in your target folder with the files split into folders (example 5 files per folder):

<p align="left">
  <!-- <a href="https://npmjs.com/package/retonio"><img src="https://badgen.net/npm/v/retonio" alt="npm package"></a> -->
  <img src="https://ik.imagekit.io/vrfrbvdn0j/sddev/before_after.png?ik-sdk-version=javascript-1.4.3&updatedAt=1643641225915" alt="GoScission before after">
</p>

## Why

I run frequently into the case, that I need to upload files to a service, which accepts only 50 by a time. By that I can upload folder by folder instead of checking by hand, which are the next 50 files.

## Thanks for using