# Write title here

## Overview

Write overview here

## Technologies Used

- **go-playground/validator** ([https://github.com/go-playground/validator](https://github.com/go-playground/validator)): A Go library for data validation and sanitization.

These technologies provide a robust and efficient foundation for building modern GraphQL APIs in Go.

## Running the Application

1. Clone the repository:

   ```bash
    git clone github.com/Xta-Ama/Mservice.git
   ```

2. Navigate to the project directory:

   ```bash
    cd Mservice
   ```

3. Generate the swagger docs:

   ```bash
    make swagger-generate
   ```

4. Run the application in watch mode:

   ```bash
    make watch
   ```

## Contributing

Contributions are welcome! To contribute to this project, follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature/your-feature`).
6. Create a new Pull Request.

## Commit Convention

Before you create a Pull Request, please check whether your commits comply with
the commit conventions used in this repository.

When you create a commit we kindly ask you to follow the convention
`category(scope or module): message` in your commit message while using one of
the following categories:

- `feat / feature`: all changes that introduce completely new code or new
  features
- `fix`: changes that fix a bug (ideally you will additionally reference an
  issue if present)
- `refactor`: any code related change that is not a fix nor a feature
- `docs`: changing existing or creating new documentation (i.e. README, docs for
  usage of a lib or cli usage)
- `build`: all changes regarding the build of the software, changes to
  dependencies or the addition of new dependencies
- `test`: all changes regarding tests (adding new tests or changing existing
  ones)
- `ci`: all changes regarding the configuration of continuous integration (i.e.
  github actions, ci system)
- `chore`: all changes to the repository that do not fit into any of the above
  categories

  e.g. `feat(authentication): add validation step to validate auth payload`

If you are interested in the detailed specification you can visit
<https://www.conventionalcommits.org/> or check out the
[Angular Commit Message Guidelines](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines).

## License

Include license information if applicable.

## Contact

For any inquiries or support, please contact [your-email](mailto:your-email@example.com).
