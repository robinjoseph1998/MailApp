# Page 2

## Example .md File

Design and Development tips in your inbox. Every weekday.

ads via Carbon

Advertising sustains the DA. Ads are hidden for members. Join today

On this page

README format

README sections

Project name and introduction

Table of contents (TOC)

Requirements

Recommended modules

Installation

Configuration

Troubleshooting & FAQ

Maintainers

Changes from README.txt

Advanced help module

Sample README

Documenting your project

Help text standards

Module documentation guidelines

Project page template

README Template

Working with Drupal's Help system

Using simplytest.me as a project demo

Help Topic Standards

README.md template

Last updated on 13 March 2024

The Drupal community recommends using the Markdown format to create a README.md file. Please note that the Drupal Coding Standards have not yet been updated to allow Markdown files, conversion should only be undertaken at the request of an existing project maintainer.

For a quick introduction to Markdown, see Markdown Guide's Basic Syntax or GitLab Flavored Markdown (GLFM) for a more comprehensive run-down. Please also review the Drupal Coding Standards for Markdown files prior to making changes.

README documentation can include more Markdown-features than those used in the example snippets in this template. For instance, Markdown let you create hyperlinks that can be clicked when the README.md is rendered in GitLab or by a compatible program (e.g. Advanced Help).

You can generate a boilerplate README.md with the following command: drush generate readme. The generated text does not follow these guidelines exactly, but may serve as a good starting point.

README format

Drupal recommends the following README formatting:

Headings capitalized with an initial capital, following standard English sentence rules

Headings prefixed with #/##/### to indicate level of heading (h1/h2/h3) followed by a blank line

Project name is the first line of the document, and only level one heading (#)

Two lines prior to ##/### headings

No leading or trailing spaces

Bulleted lists denoted by dashes (-)

Ordered lists use "1", for easier updates and to avoid errors (see Configuration)

Nested lists indented with 4 spaces

Links should have a meaningful link text, for example:

\[Drupal]\(https://www.drupal.org/) (i.e. not just the URL)

Text manually word-wrapped within around 80 cols

README sections

Drupal recommends the following README sections:

Project name and introduction (required)

Table of contents (optional)

Requirements (required)

Recommended modules (optional)

Installation (required, unless a separate INSTALL.md is provided)

Configuration (required)

Troubleshooting & FAQ (optional)

Maintainers (optional)

Project name and introduction

Start the README.md with the project name, and an introduction to the project. The project name is the only level one heading in the document. This must be the first line of the document and must be followed by one blank line.

The introduction summarizes the purpose and function of the project, and should be concise (a brief paragraph or two). This introduction may be the same as the first paragraph on the project page.

This section should include a link to the project page and issue queue. If the project is a sandbox, these links should go to the sandbox until promotion.

## Administration Menu

The Administration Menu module displays the entire administrative menu tree

(and most local tasks) in a drop-down menu, providing administrators one- or

two-click access to most pages.

For a full description of the module, visit the

\[project page]\(https://www.drupal.org/project/admin\_menu).

Submit bug reports and feature suggestions, or track changes in the

\[issue queue]\(https://www.drupal.org/project/issues/admin\_menu).

Table of contents (TOC)

TOCs are optional but appreciated for lengthy README files.

### Table of contents

* Requirements
  * Recommended modules
  * Installation
  * Configuration
  * Troubleshooting
  * FAQ
  * Maintainers

Requirements

The requirements section describes whether this project requires anything outside of Drupal core to work (modules, libraries, etc). List all requirements here, including those that follow indirectly from another module, etc. The idea is to inform the users about what is required, so that everything they need can be procured and included in advance of attempting to install the module. If there are no requirements, write "No special requirements".

### Requirements

This module requires the following modules:

* \[Views]\(https://www.drupal.org/project/views)
  * \[Panels]\(https://www.drupal.org/project/panels)

### Requirements

This module requires no modules outside of Drupal core.

Recommended modules

The optional recommended modules section lists modules that are not required, but that may enhance the usefulness or user experience of your project. Make sure to describe the benefits of enabling these modules.

### Recommended modules

\[Markdown filter]\(https://www.drupal.org/project/markdown): When enabled,

display of the project's README.md help will be rendered with markdown.

Installation

The installation section describes how to install the module. However, if the steps to install the module follow the standard instructions for Drupal 9, Drupal 7, Drupal 6/5, or a theme, don't reinvent the wheel — simply provide a link and explain in detail any steps that may diverge from these steps. Take special note of Drush integrations. In a case where many Drush commands are added, consider adding a section for Drush.

Consider replacing this section with a standalone INSTALL.md file if your installation instructions are especially complex.
