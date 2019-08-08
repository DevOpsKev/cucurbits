<img src="https://github.com/TideKev/cucurbits/tree/master/images/cucmbers.png">

# Cucurbits

A DevOps Native HDD framework written in Ruby. Cucurbits uses Gherkin like syntax to describe and test product hypotheses.

### Some Links on HDD:

https://www.thoughtworks.com/de/insights/blog/how-implement-hypothesis-driven-development

https://opensource.com/article/19/6/why-hypothesis-driven-development-devops

https://hackerchick.com/hypothesis-driven-development/

### Why Cucurbits?

DevOps needs a lightweight framework that will allow for executable specification of **Product Hypothesis**. You can add **Cucurbits** to your toolchain and achieve complete automation from **Idea** to **Implementation**.

**Cucurbits** is an aid to communication between **Product Development** and other contributors involved in the product delivery lifecycle.

**Cucurbits** operate on a **Belief System** and uses Gherkin like syntax to describe **Beliefs** stored in `.belief` files.

Within a **Belief** file, we describe the basis or bases for our beliefs and methods to test the bases of our beliefs. 

When we test the basis of our beliefs, we may receive **Confirmation** of our beliefs when those tests pass. When we test the basis of our beliefs, and we are proven comprehensively wrong, we will receive a **Revelation**. The example below demonstrates why we believe DevOps needs this framework, it works like this:

````
Cohort : People
  Criteria : Homo sapiens.
  Criteria : Have a Github account.
  Criteria : Live on this planet.

Belief : The world needs a lightweight HDD framework to describe and test product hypotheses.
  Basis : Executable specification reduces ambiguity in the DevOps lifecycle.
    Test : Release an opensource HDD framework.
      Criteria : Allow description of beliefs in natural language for non-technical people.
      Criteria : Beliefs must be captured as an Executable Specification.
      Criteria : Developed in familiar tools and languages for engineers.
      Criteria : Infrastructure as Code deployable.
      Criteria : Capable of managing long-running tests.
      AttentionSpan : 14 days
  Comfirmation : [>2000](People) forked us on Github AND [>100](People) contributed to the codebase.
  Revelation : [<10](People) forked us on Github. (People) just aren't interested in working smarter.
````

### How To Guide:

1. Install **Cucurbits**.

2. Capture your **Beliefs** in a `.belief` file, eg: `the-world-needs-a-lightweight-hdd-framework.belief`.

3. Run `cucurbits init` in the directory containing your `.belief` file.

4. **Cucurbits** will generate two subdirectories `./cucerbits-tests` and `./cucerbits-service`.

5. Within the `./cucerbits-tests` directory will be a file named `cucurbits-test-1.rb` containing code stubs to implement your live tests.

6. Each **Belief** to be **Tested** in the `cucurbits-test-1.rb` will be issued with a **Public Key.** You will need this when calling the **Cucerbits Service**.

7. Implement and deploy code to test your beliefs.

8. Install **Terraform**.

9. In the `./cucerbits-service` directory, you will find a **Terraform** example configuration file to deploy **Cucerbits Service** named `cucurbits.tf`.

10. Follow the instructions commented within the `cucurbits.tf` file to configure deployment for your environment.

11. From within the `./cucerbits-service` directory execute the commands `terraform init` and `terraform apply`.

12. You now have a running **Cucerbits Service** in your environment.

13. Subscribe to **Cucurbits Service** to receive either **'Confirmation'** of your **Beliefs** or a **'Revelation.'**

14. Check out the **Cucurbits Service UI** to browse your **Belief System**.

15. Additional notes:
    1. **Cucurbits Service** will only start testing **Beliefs** once it has received notification all **Test Criteria** have been deployed.
    2. If a given **Belief's AttentionSpan** expires the **Cucurbits Service** will stop testing that **Belief**.
    3. Each **Belief** is protected by a unique key pair, you will need the **Public Key** when subscribing or publishing to a **Belief**.
