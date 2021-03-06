import tega.tree
import tega.subscriber

import os
import subprocess

class Deployment(tega.subscriber.PlugIn):
    '''
    Deployment of NLAN services
    '''


    def __init__(self):
        super().__init__()
        try:
            self.script = os.environ['SETUP_SCRIPT']
        except:
            self.script = 'setup.sh'
        self.scriptdir = os.path.join(os.environ['GOPATH'], 'src/github.com/araobp/nlan/')
        plugins = tega.tree.Cont('plugins')
        with self.tx() as t:
            plugins.deploy = self.func(self._deploy)  # Attached to plugins.deploy
            t.put(plugins.deploy, ephemeral=True)

    def on_notify(self, notifications):
        pass

    def on_message(self, channel, tega_id, message):
        pass

    def _deploy(self):
        '''
        Deployment
        '''
        os.chdir(self.scriptdir)
        routers = self.get("ip").keys()
        args = [os.path.join(self.scriptdir, self.script)]
        args.extend(routers)
        self.process = subprocess.Popen(args, preexec_fn=os.setsid)
