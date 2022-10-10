import React from 'react';

import mm_constants from 'mattermost-redux/constants/general';

import BaseCard from 'components/card/base';
import IconButton from 'components/buttons/iconButton';
import LabelValuePair from 'components/labelValuePair';
import SVGWrapper from 'components/svgWrapper';

import plugin_constants from 'plugin_constants';
// const boards_icon = require('../../../../../assets/boards_icon.svg')
import './styles.scss';
import { boards } from 'plugin_constants/common';

type SubscriptionCardProps = {
    handleDeleteSubscrption: (subscriptionDetails: SubscriptionDetails) => void
    subscriptionDetails: SubscriptionDetails
}

const SubscriptionCard = ({handleDeleteSubscrption, subscriptionDetails: {channelType, eventType, serviceType, channelName, createdBy}, subscriptionDetails}: SubscriptionCardProps) => (
    <BaseCard>
        <div>
            <div className='d-flex justify-content-between align-items-center mb-1'>
                <div className="d-flex">
                    <span>
                        <SVGWrapper
                            width={20}
                            height={20}
                            viewBox=' 0 0 20 20'
                        >
                            {serviceType==boards ? plugin_constants.SVGIcons.boards : plugin_constants.SVGIcons.repos}
                        </SVGWrapper>
                    </span>
                    <p className={`ml-1 color-${serviceType}`}>{serviceType.charAt(0).toUpperCase() + serviceType.slice(1)}</p>
                </div>
                <div className='button-wrapper'>
                    <IconButton
                        tooltipText='Delete subscription'
                        iconClassName='fa fa-trash-o'
                        extraClass='delete-button'
                        onClick={() => handleDeleteSubscrption(subscriptionDetails)}
                    />
                </div>
            </div>
            <div className='project-details'>
                <LabelValuePair
                    label={
                        <SVGWrapper
                            width={12}
                            height={12}
                            viewBox='0 0 10 10'
                        >
                            {plugin_constants.SVGIcons.workEvent}
                        </SVGWrapper>
                    }
                    labelExtraClassName='margin-left-5'
                    value={plugin_constants.common.eventTypeMap[eventType as EventType] ?? ''}
                />
                <LabelValuePair
                    labelIconClassName={`icon ${channelType === mm_constants.PRIVATE_CHANNEL ? 'icon-lock-outline' : 'icon-globe'} icon-label`}
                    value={channelName}
                />
                <LabelValuePair
                    labelIconClassName={'icon icon-account-outline icon-label'}
                    value={`Subscription created by ${createdBy}`}
                />
            </div>
        </div>
    </BaseCard>
);

export default SubscriptionCard;
