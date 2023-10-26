import os
import subprocess

from templateframework.metadata import Metadata


def run(metadata: Metadata = None):
    subprocess.run("go mod tidy", shell=True, cwd=str(metadata.target_path)+"/app")
    return metadata
