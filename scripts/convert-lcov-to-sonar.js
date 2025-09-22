#!/usr/bin/env node

/**
 * Convert LCOV format to SonarQube Generic Coverage XML format
 */

const fs = require('fs');
const path = require('path');

function parseLcov(lcovContent) {
  const lines = lcovContent.split('\n');
  const files = [];
  let currentFile = null;

  for (const line of lines) {
    if (line.startsWith('SF:')) {
      if (currentFile) {
        files.push(currentFile);
      }
      currentFile = {
        path: line.substring(3),
        lines: []
      };
    } else if (line.startsWith('DA:') && currentFile) {
      const parts = line.substring(3).split(',');
      const lineNumber = parseInt(parts[0]);
      const hits = parseInt(parts[1]);
      currentFile.lines.push({ lineNumber, hits });
    } else if (line === 'end_of_record' && currentFile) {
      files.push(currentFile);
      currentFile = null;
    }
  }

  return files;
}

function generateSonarXml(files) {
  let xml = '<?xml version="1.0" encoding="UTF-8"?>\n';
  xml += '<coverage version="1">\n';

  for (const file of files) {
    xml += `  <file path="${file.path}">\n`;

    for (const line of file.lines) {
      const covered = line.hits > 0 ? 'true' : 'false';
      xml += `    <lineToCover lineNumber="${line.lineNumber}" covered="${covered}"/>\n`;
    }

    xml += '  </file>\n';
  }

  xml += '</coverage>\n';
  return xml;
}

function main() {
  const lcovFile = process.argv[2] || 'coverage.lcov';
  const outputFile = process.argv[3] || 'coverage-sonar.xml';

  if (!fs.existsSync(lcovFile)) {
    console.error(`LCOV file not found: ${lcovFile}`);
    process.exit(1);
  }

  const lcovContent = fs.readFileSync(lcovFile, 'utf8');
  const files = parseLcov(lcovContent);
  const sonarXml = generateSonarXml(files);

  fs.writeFileSync(outputFile, sonarXml);
  console.log(`Converted ${lcovFile} to ${outputFile}`);
  console.log(`Files processed: ${files.length}`);
}

if (require.main === module) {
  main();
}

module.exports = { parseLcov, generateSonarXml };