<img src="https://media.giphy.com/media/3oEduQAsYcJKQH2XsI/giphy.gif" >

# Cucerbits

A DevOps native HDD framework written in Ruby, Cucurbits uses Gherkin like syntax to describe and test product hypotheses. 

### Some Links on HDD:

https://opensource.com/article/19/6/why-hypothesis-driven-development-devops

https://www.thoughtworks.com/de/insights/blog/how-implement-hypothesis-driven-development

https://hackerchick.com/hypothesis-driven-development/

### Why Cucerbits?

DevOps needs a lightweight framework that will allow for executable specification of 'Product Hypothesis'. You can add Cucerbits to your toolchain and achieve complete automation from 'Idea' to 'Implementation'.

Cucerbits will aid copmmunication between Product Development and other contributors involved in the product delivery lifecycle. 

Cucerbits operates on a 'Belief System' and uses Gherkin like systnax to describe 'Beliefs' that are then stored in a '.belief' file. 

Within a belief file we describe the basis for those beliefs and methods to test our beliefs. When we test the basis of our beliefs we will either recieve 'Confirmation' of our beliefs when those tests pass. When we test the basis of our beliefs and we are proven comprehensively wrong we will receive a revelation. It works like this:


````
Belief : The world needs a lightweight HDD framework to describe and test product hypotheses.
  Basis : Executable specification reduces ambiguity in the DevOps lifecycle.
    Test : Release an opensource HDD framework.
      Criteria : Allow description of beliefs in natural language for non-technical people.
      Criteria : Beliefs must be captured as an Executable Specification.
      Criteria : Developed in familiar tools and languages for engineers.
      Criteria : Infrastructure as Code deployable.
      Criteria : Capable of managing long-running tests.
      AttentionSpan : 14 days
  Comfirmation : People use the framework, fork us on github and contribute to opensource development.
  Revelation : Noone is interested, we don't recieve a single fork. People actually enjoy ambiguity.
````



