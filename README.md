# databahn-assignment

## Problem statement
1. Develop an API endpoint named "/load/{director}/{template_file}".
2. Include optional query parameters, such as "count," for enhanced flexibility.
3. The core functionality of the "load" operation is to take a template file and replace template variables with actual values. For instance, when encountering a "timestamp" variable, it should be replaced with the current timestamp.
4. The processed data must be pushed to a specified destination, such as Kafka or a file storage system.
5. The system should be optimized to efficiently handle a substantial volume of data. Specifically, it should be capable of processing at least 50,000 records effectively.
