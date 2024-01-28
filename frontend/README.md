# Backend

This backend is a critical component that handles authentication through GitHub OAuth, and integrating with GitHub API.

## Tech Stack

### Next.JS 13
Next.js typically introduces improvements in performance, developer experience, and additional features with each release. Version upgrades often include enhancements to build times, optimization of server-side rendering, and updates to the underlying dependencies. Additionally, new versions might bring improvements to the development server, enhanced support for serverless functions, and updates to the Next.js API. Always check the official Next.js release notes for the most accurate and up-to-date information on the advantages of specific versions.

### Tailwind CSS
Tailwind CSS makes styling websites super easy. It gives you ready-to-use classes that help you design without having to write a lot of custom CSS. This makes it quick to create consistent and responsive designs, and you can easily tweak styles as needed. Tailwind simplifies the whole process, making it efficient for building and adjusting the look of your website without the hassle of writing a ton of CSS code.

### React Context
React Context makes it easy to share information between different parts of a React app without passing it through many components. This helps keep the code clean and avoids unnecessary complications. It's handy for managing things like themes or user preferences across the whole application, making state management in React simpler.

### No External REST API Client
We connect with the backend using the built-in libraries and the Next.js API, removing the necessity for extra external REST API client libraries. This simplifies the frontend by reducing its complexity.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/B0urG3ois/pixel8labs-test.git

2. Install Node.js and NPM:
   ```bash
   https://docs.npmjs.com/downloading-and-installing-node-js-and-npm

3. Run application:
   ```bash
   cd /frontend
   npm install
   npm run dev

4. Application Start:
   ```bash
   Application start on port 3000, and open your browser
   http://localhost:3000

## Documentation

- Figma Design
  ```bash
  https://www.figma.com/file/fLiLQfjSF6X7pEfHli2Lwh/Fullstack-Engineer-Test-Case?type=design&node-id=0%3A1&mode=design&t=RfULQB2MF956TxTT-1

- Frontend Public URL
  ```bash
  https://pixel8labs-assignment.vercel.app/

----------------------------------------------------------------
Please feel free to examine the code and make any required modifications or improvements to customize it for your particular needs.
