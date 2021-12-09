resource "aws_security_group" "allow_ssh" {
  name        = "allow_all"
  description = "Allow inbound SSH traffic from any IP"
  vpc_id      = "VPC-ID"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags {
    Name = "Allow SSH"
  }
}

resource "aws_security_group" "allow_postgres" {
  name        = "allow-postgres"
  description = "Allow Postgres traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description      = "Postgres from VPC"
    from_port        = 5432
    to_port          = 5432
    protocol         = "tcp"
    cidr_blocks      = [aws_vpc.main.cidr_block]
    ipv6_cidr_blocks = [aws_vpc.main.ipv6_cidr_block]
  }

  tags = {
    Name = "allow_intra_vpc_postgres"
  }
}

resource "aws_security_group" "http_service" {
  name        = "dev-service"
  description = "Allow HTTP in, all outbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "http in"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "dev_service"
  }
}

resource "aws_security_group" "search_service" {
  name        = "dev-search"
  description = "Allow TCP search traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "search in"
    from_port   = 9200
    to_port     = 9200
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "dev_search"
  }
}
