package ksyaml

import (
	"strings"

	"testing"
)

type TestCase struct {
	input    string
	expected string
	name     string
}

func TestConvert(t *testing.T) {

	cases := []TestCase{
		{
			name: "Simple root key value pair with comment",
			input: `
# in atas
version: "3.8" # inline gan
#ini bawah
num: 1
num lagi: 1.10
boolean: false
kalo ini string: ok gan
# ini comment
# asd

# asd`,
			expected: `
# in atas
version: "3.8" # inline gan
#ini bawah
num: 1
num lagi: 1.10
boolean: false
kalo ini string: "ok gan"
# ini comment
# asd
# asd`,
		}, {
			name: "Nested Object with Comment",
			input: `
services: # ok
  db: # ini comment lain
  # comment with indent
    container_name: scelefeed-db
    image: postgres:13.3-alpine
    volumes:
      - db-data:/var/lib/postgresql/data
    restart: always
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: ${DB_NAME} # inline comment habis koma
    networks:
      - scelefeed
    ports:
      - "${DB_PORT}:5432"
}`,
			expected: `
services: { # ok
  db: { # ini comment lain
    # comment with indent
    container_name: "scelefeed-db",
    image: "postgres:13.3-alpine",
    volumes: [
      "db-data:/var/lib/postgresql/data"
    ],
    restart: "always",
    environment: {
      POSTGRES_USER: "${DB_USER}",
      POSTGRES_PASSWORD: "${DB_PASSWORD}",
      POSTGRES_DB: "${DB_NAME}" # inline comment habis koma
    },
    networks: [
      "scelefeed"
    ],
    ports: [
      "${DB_PORT}:5432"
    ]
  }
}`,
		},
		{
			name: "Nested object with multiple children and comment",
			input: `
services: # ok
  db: # ini comment lain
    # comment with indent
    container_name: scelefeed-db
    image: postgres:13.3-alpine
    volumes:
      - db-data:/var/lib/postgresql/data
    restart: always
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: ${DB_NAME} # inline comment habis koma
    networks:
      - scelefeed
    ports:
      - "${DB_PORT}:5432"

  server:
    container_name: server-scelefeed
    build:
      context: .
      dockerfile: Dockerfile # inline comment habis koma
    ports:
      - "${SERVER_PORT}:${SERVER_PORT}"
    depends_on:
      - db
    environment:
      - DB_HOST=db # inline comment habis koma di array
      - DB_PORT=5432 # inline comment habis koma di array
    networks:
      - scelefeed
    command:
      - "/app/main"`,
			expected: `
services: { # ok
 db: { # ini comment lain
  # comment with indent
  container_name: "scelefeed-db",
  image: "postgres:13.3-alpine",
  volumes: [
   "db-data:/var/lib/postgresql/data"
  ],
  restart: "always",
  environment: {
   POSTGRES_USER: "${DB_USER}",
   POSTGRES_PASSWORD: "${DB_PASSWORD}",
   POSTGRES_DB: "${DB_NAME}", # inline comment habis koma
  },
  networks: [
   "scelefeed"
  ],
  ports: [
   "${DB_PORT}:5432"
  ],
 },
 server: {
  container_name: "server-scelefeed",
  build: {
   context: ".",
   dockerfile: "Dockerfile", # inline comment habis koma
  },
  ports: [
   "${SERVER_PORT}:${SERVER_PORT}"
  ],
  depends_on: [
   "db"
  ],
  environment: [
   "DB_HOST=db", # inline comment habis koma di array
   "DB_PORT=5432" # inline comment habis koma di array
  ],
  networks: [
   "scelefeed"
  ],
  command: [
   "/app/main"
  ],
 },
}`,
		},
}
	for _, tc := range cases {

		in := strings.TrimSpace(tc.input)
		exout := strings.TrimSpace(tc.expected)

		out, err := Convert(in)
		if err != nil {
			t.Errorf("Test case '%s' returns an error: %s", tc.name, err)
		}
		if out != exout {
			t.Errorf("\nTest case '%s' returns an unexpected output:\n %s\n\nexpected:\n %s", tc.name, out, exout)
		}
	}
}
