<img src="https://media.giphy.com/media/3oEduQAsYcJKQH2XsI/giphy.gif" >

# Cucerbits

A DevOps Native HDD framework written in Ruby. Cucurbits uses Gherkin like syntax to describe and test product hypotheses.

### Some Links on HDD:

https://opensource.com/article/19/6/why-hypothesis-driven-development-devops

https://www.thoughtworks.com/de/insights/blog/how-implement-hypothesis-driven-development

https://hackerchick.com/hypothesis-driven-development/

### Why Cucerbits?

DevOps needs a lightweight framework that will allow for executable specification of 'Product Hypothesis.' You can add Cucurbits to your toolchain and achieve complete automation from 'Idea' to 'Implementation.'

Cucurbits will aid communication between Product Development and other contributors involved in the product delivery lifecycle.
Cucurbits operate on a 'Belief System' and uses Gherkin like syntax to describe 'Beliefs' stored in '.belief' files.

Within a belief file, we describe the basis for our beliefs and methods to test our beliefs. 

When we test the basis of our beliefs, we may receive 'Confirmation' of our beliefs when those tests pass. When we test the basis of our beliefs, and we are proven comprehensively wrong, we will receive a revelation. It works like this:

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

### How To Guide:

1. Install Cucurbits.

2. Capture your 'Beliefs' in a .belief file.

3. Run `cucurbits init` in the directory containing your .belief file.

4. Cucurbits will generate two subdirectories ./cucerbits-tests and ./cucerbits-server.

5. Within the `./cucerbits-tests` directory will be a file named `cucurbits-test-1.rb` containing code stubs to implement your live tests.

6. Each 'Belief' to be 'Tested' in the cucurbits-test-1.rb will be issued with a 'Public Key.' You will need this when calling the Cucerbits server.

7. Implement and deploy code to test your beliefs.

8. Install 'Terraform'.

9. In the ./cucerbits-server directory, you will find a Terraform example configuration file to deploy Cucerbits server named cucurbits.tf.

10. Follow the instructions commented within the cucurbits.tf file to configure deployment for your environment.

11. From within the `./cucerbits-server` directory execute the commands terraform init and terraform apply.

12. You now have a running Cucerbits server in your environment.

13. Subscribe to Cucurbits to receive either 'Confirmation' of your beliefs or a 'Revelation.'

14. Check out the Cucurbits Server UI to browse your Belief System.

15. Additional notes:
    1. Cucurbits Server will only start testing beliefs once it has received notification 'Test Criteria' have been deployed.
    2. If a given Belief's AttentionSpan expires Cucurbits will stop testing that Belief.
    3. Each Belief is protected by a unique key pair, you will need the 'Public Key' when subscribing or publishing to a Belief.
